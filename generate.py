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
    with open(name + '.go', 'w') as f:
        f.write('package ' + package + '\n')
        f.write(''.join(list))

templates = dict()

all_operations = []
all_conversions = []

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

operations = ['Add', 'Subtract', 'Multiply', 'Divide']

for interface, underlying in interfaces.items():
    for op in operations:
        all_operations.append(template('operation', interface=interface, op=op,
                              pascal_underlying=underlying))
    for camel, pascal in types:
        all_conversions.append(template('conversion', interface=interface,
                               pascal_underlying=underlying, pascal_type=pascal,
                               camel_type=camel))

save_list('operations', all_operations)
save_list('conversions', all_conversions)
