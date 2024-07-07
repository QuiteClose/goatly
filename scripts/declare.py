#/bin/env python3
'''Usage: python3 declare.py <PACKAGE> <UNLESS PACKAGE PATH> <DECLARE PACKAGE PATH>

Reads an internal/unless package from stdin and writes the corresponding
declare package to stdout, where <PACKAGE> is one of:

    *   assert
    *   expect
'''

import os.path
import sys

###############################################################################

PACKAGES = {
    'assert': ('AssertFailed', 't.Fatalf'),
    'expect': ('ExpectFailed', 't.Errorf'),
}

PACKAGE_HEADER = '''
// Re-implementation of the internal/unless package generated by declare.py
// Each function calling {testing_call} with an {error_type} message if the
// condition is not met.
package {package_name}

import (
	"regexp"
	"testing"

	"github.com/quiteclose/goatly/internal/unless"
	"github.com/quiteclose/goatly/pkg/run"
)

///////////////////////////////////////////////////////////////////////////////
'''.strip()

DECLARE_FUNCTION = '''
{comment}func {declare_name}(t *testing.T, {args}, message string) bool {{
	return unless.{unless_name}({arg_names}, func(s string) {{
		{testing_call}("{error_type}: %s\\n%s", message, s)
	}})
}}'''

###############################################################################

# Example Function:
# func Equal(a, b interface{}, callback func(string)) {
def parse_function(line, comment, error_type, testing_call):
    name = line.split('func ')[1].split('(')[0]
    args = line.split('(')[1].split(', callback func')[0]
    arg_names = ', '.join([arg.split(' ')[0] for arg in args.split(', ')])
    return DECLARE_FUNCTION.format(
        comment=comment.replace('call the callback', f'call {testing_call}'),
        declare_name=name,
        unless_name=name,
        args=args,
        arg_names=arg_names,
        testing_call=testing_call,
        error_type=error_type,
    )


def read_functions(stream):
    comment_block = []
    for line in stream:
        if line.startswith('// '):
            comment_block.append(line)
        else:
            if line.startswith('func '):
                yield line.strip(), ''.join(comment_block)
            else:
                comment_block = []


def parse_package(package_name, error_type, testing_call, *source_paths):
    result = []
    functions = []
    result.append(PACKAGE_HEADER.format(
        package_name=package_name,
        error_type=error_type,
        testing_call=testing_call,
    ))
    for source_path in source_paths:
        with open(source_path) as reader:
            for name, comment in read_functions(reader):
                functions.append((name, comment))
    for name, comment in sorted(functions):
        result.append(parse_function(name, comment, error_type, testing_call))
    return '\n'.join(result)


def run(package_name, error_type, testing_call, target_path, *source_paths):
    with open(target_path, 'w') as writer:
        print(parse_package(package_name, error_type, testing_call, *source_paths), file=writer)


if __name__ == '__main__':
    if len(sys.argv) != 4 or sys.argv[1] not in PACKAGES:
        print(__doc__)
        sys.exit(1)
    result = []
    package_name = sys.argv[1]
    unless_path = sys.argv[2]
    declare_path = sys.argv[3]
    error_type, testing_call = PACKAGES[package_name]
    target_path = os.path.join(declare_path, package_name+'.go')
    source_paths = [
        os.path.join(unless_path, file_name) for file_name in
        ['unless.go', 'run.go']
    ]
    run(package_name, error_type, testing_call, target_path, *source_paths)

