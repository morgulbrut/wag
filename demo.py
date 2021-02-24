#!/usr/bin/python3
import os

teststring = '''ABCDEFGHIJKLMNOPQRSTUVWXYZ 1234567890 ()!?:;.,'''

for i in range(22):
    fn = "font{}".format(i+1)
    print("--------------------------------------")
    print(fn)
    print("--------------------------------------")

    cmd = "wag.exe -t " + '"' + fn.upper() + ": " + teststring + \
        '"' + " -f "+fn + " -o example/" + fn

    os.system(cmd)

for i in range(22):
    fn = "font{}".format(i+1)
    line = "!["+fn+"](example/"+fn+".png)"
    print(line)
