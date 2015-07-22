#!/usr/bin/env python3

package = 'numeric'

def read_template(name):
    with open('template_' + name, 'r') as f:
        return ''.join(f.readlines())

def template(name, **kwargs):
    if not name in templates:
        templates[name] = read_template(name)
    str = templates[name]
    for k, v in kwargs.items():
        str = str.replace('<' + k.upper() + '>', v)
    return str

def save_list(name, list):
    with open('generated_' + name + '.go', 'w') as f:
        f.write('package ' + package + '\n')
        f.write(''.join(list))

templates = dict()

types = [
    ('uint32', 'UInt32'),
    ('int32', 'Int32'),
    ('uint64', 'UInt64'),
    ('int64', 'Int64'),
    ('float32', 'Float32'),
    ('float64', 'Float64')
]

interfaces = {
    'Integer': 'Int64',
    'UInteger': 'UInt64',
    'Float': 'Float64'
}

operations = {
    'Add': '+',
    'Subtract': '-',
    'Multiply': '*',
    'Divide': '/',
}

all_operations = []
all_conversions = []
all_unary_operations = []

for interface, underlying in interfaces.items():
    for op_name, op_symbol in operations.items():
        all_operations.append(template('operation', interface=interface,
                              op_name=op_name, op_symbol=op_symbol,
                              pascal_underlying=underlying, return_type='Numeric'))
    for camel, pascal in types:
        all_conversions.append(template('conversion', interface=interface,
                               pascal_underlying=underlying, pascal_type=pascal,
                               camel_type=camel))
    all_unary_operations.append(template('unary_operations', interface=interface))

save_list('operations', all_operations)
save_list('conversions', all_conversions)
save_list('unary_operations', all_unary_operations)
