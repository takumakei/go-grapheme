#!/usr/bin/env python3

import fire
import re
import sys


def main():
    """ main """

    print("package grapheme")
    print()
    print("var testCases = []testCase {")

    d = '÷'
    m = '×'
    h = '[0-9a-fA-F]+'

    r = re.compile('\s*{d}\s*({h}(?:\s*(?:{d}|{m})\s*{h})+)\s*{d}\s*#\s*(.*)'.format(d=d, m=m, h=h))
    s = re.compile(d)
    t = re.compile(m)
    for line in readline(sys.stdin):
        m = r.match(line)
        if not m:
            continue
        x = m.group(1)
        a = [[int(j, 16) for j in t.split(i.strip())] for i in s.split(x)]
        print('{{"{} # {}",'.format(x, m.group(2)), end='')
        print('"{}",'.format("".join([c_str(to_utf8(j)) for i in a for j in i])), end='')
        print('[]string{{{}}}}},'.format(', '.join(['"{}"'.format(''.join([c_str(to_utf8(j)) for j in i])) for i in a])))

    print("}")


def to_utf8(codepoint):
    if codepoint <= 0x007f:
        return [codepoint]
    if codepoint <= 0x07ff:
        return [0xc0 | (0x1f & (codepoint >>  6)), 0x80 | (0x3f & codepoint)]
    if codepoint <= 0xffff:
        return [0xe0 | (0x0f & (codepoint >> 12)), 0x80 | (0x3f & (codepoint >>  6)), 0x80 | (0x3f & codepoint)]
    if codepoint <= 0x10ffff:
        return [0xf0 | (0x07 & (codepoint >> 18)), 0x80 | (0x3f & (codepoint >> 12)), 0x80 | (0x3f & (codepoint >> 6)), 0x80 | (0x3f & codepoint)]
    raise "BAD CODEPOINT {:04x}".format(codepoint)


def c_str(a):
    return "".join(['\\x{:02x}'.format(i) for i in a])


def readline(reader):
    """ readline generator """
    a = reader.readline()
    while a:
        yield a.rstrip("\n")
        a = reader.readline()


if __name__ == '__main__':
    fire.Fire()

# vim: fileencoding=utf-8
