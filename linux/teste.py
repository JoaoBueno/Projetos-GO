from ctypes import *
import sys

if len(sys.argv) < 2:
    print("missing argument", file=sys.stderr)
    sys.exit(1)

path = sys.argv[1]

lib = cdll.LoadLibrary("libimgutil.so")

get_image_size = lib.ImgutilGetImageSize
get_image_size.argtypes = [c_char_p, POINTER(c_ulonglong), POINTER(c_ulonglong)]
get_image_size.restype = c_char_p

w = c_ulonglong()
h = c_ulonglong()

err = get_image_size(path.encode('utf-8'), byref(w), byref(h))
if err != None:
    print("error: %s" % (str(err, encoding='utf-8')))
    sys.exit(1)

print("%s: %dx%d" % (path, w.value, h.value))
