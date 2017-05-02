#!/usr/bin/env python3

import fire
import re
import sys


def main():
    """ main """
    a = read_trie(sys.stdin)

    print("package grapheme")
    print()
    print("//go:generate stringer -type=Prop $GOFILE")
    print()
    print("type Prop int")
    print()
    print("const (")
    l = labels(a)
    i = l[0]
    print("  {} Prop = iota".format(*i))
    for i in l[1:]:
        print("  {:18s} // {:3d} rules found".format(*i))
    print(")")
    print()

    a.sort(key=lambda c: c[0])

    b = regularize(a)
    b.append((0x10FFFF, "OutOfRange"))

    c = make_tree(b)
    print("func Property(r rune) Prop {")
    print_tree(c)
    print("}")
    print()


def property_test():
    a = read_trie(sys.stdin)
    a.sort(key=lambda c: c[0])
    b = regularize(a)
    b.append((0x10FFFF, "OutOfRange"))
    print("package grapheme")
    print()
    print("var keys = [...]rune{")
    for i in b:
        print("  0x{:04x},".format(i[0]))
    print("}")
    print("var klen = len(keys)")
    print()
    print("var vals = [...]Prop{")
    for i in b:
        print("  {},".format(i[1]))
    print("}")
    print()
    print("func property(r rune) Prop {")
    print("  a, b := 0, klen")
    print("  for {")
    print("    c := b - a")
    print("    if c <= 1 {")
    print("      break")
    print("    }")
    print("    m := a + c/2")
    print("    if r < keys[m] {")
    print("      b = m")
    print("    } else {")
    print("      a = m")
    print("    }")
    print("  }")
    print("  return vals[a]")
    print("}")


def labels(tries):
    a = ["OutOfRange", "Any"]
    m = {"OutOfRange": 0, "Any": 0}
    for i in tries:
        label = i[2]
        if label not in m:
            m[label] = 1
            a.append(label)
        else:
            m[label] = m[label] + 1
    return [(i, m[i]) for i in a]


def print_tree(tree):
    print_tree_in(tree, 1)


def print_tree_in(tree, indent):
    i = "  " * indent
    a = tree[0]

    if len(tree) == 1:
        print("{}return {}".format(i, a))
    else:
        print("{}if r < 0x{:05x} {{".format(i, a))
        print_tree_in(tree[1], indent + 1)
        print("{}}} else {{ // 0x{:05x} <= r".format(i, a))
        print_tree_in(tree[2], indent + 1)
        print("{}}}".format(i));


def make_tree(ls):
    if len(ls) == 1:
        return (ls[0][1],)
    m = int(len(ls)/2)
    return (ls[m][0], make_tree(ls[:m]), make_tree(ls[m:]))


def regularize(tries):
    """[(a, A), (b, B), ...] (a以上b未満はA)というリストに変換する

    * 入力項目間のギャップはAnyで埋める
    * 連続したラベルは圧縮する
    """
    ret = []
    n = 0
    m = "Any"
    for i in tries:
        lhs, rhs, label = i

        if n > lhs:
            raise "BAD RANGE"

        if n < lhs:
            # ギャップがあるならAnyで埋めておく
            ret.append((n, "Any"))
            ret.append((lhs, label))
            m = label
        elif m != label:
            ret.append((lhs, label))
            m = label
        else:
            # ギャップなしでラベルが同じなら省略可能
            pass

        n = rhs + 1
    return ret


def print_trie(trie):
    lhs, rhs, label = trie
    if lhs == rhs:
        print("[{:05x}][-----][{}]".format(lhs, label))
    else:
        print("[{:05x}][{:05x}][{}]".format(lhs, rhs, label))


def read_trie(reader):
    """コードポイントをGrapheme Break Proertyに変換する表を、ユニコードの仕様テキストから構築する

    入力テキスト行は２つのパターンがある

    範囲指定> コードポイント始点..コードポイント終点 ; ラベル
    値指定  > コードポイント ; ラベル

    * コードポイント始点とコードポイント終点はどちらもinclusiveである。
    * 値指定の場合は始点と終点が同じ値の範囲指定に変換する

    これらをタプルのリストにする。

    [(コードポイント始点, コードポイント終点, ラベル), ...]

    * ある要素のコードポイント終点と次の要素のコードポイント始点の間に隙間があることがある
    * ある要素のコードポイント終点と次の要素のコードポイント始点の間に隙間がなくかつラベルが同じこともある(圧縮可能だ)
    """
    ret = []
    m = re.compile('^([0-9a-fA-F]+)(?:\.\.([0-9a-fA-F]+))?\s*;\s*([^\s]+)\s*')
    for line in readline(reader):
        a = m.match(line)
        if not a:
            continue

        lhs = a.group(1)
        rhs = a.group(2)
        lbl = a.group(3)

        lhs = int(lhs, 16)
        rhs = int(rhs, 16) if rhs else lhs

        ret.append((lhs, rhs, lbl))

    return ret


def readline(reader):
    """ readline generator """
    a = reader.readline()
    while a:
        yield a.rstrip("\n")
        a = reader.readline()


if __name__ == '__main__':
    fire.Fire()

# vim: fileencoding=utf-8
