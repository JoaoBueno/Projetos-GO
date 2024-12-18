from ctypes import *
import sys

if len(sys.argv) < 2:
    print("missing argument", file=sys.stderr)
    sys.exit(1)

lib = cdll.LoadLibrary("libimgutil32.so")

get_image_size = lib.ImgutilGetImageSize
get_image_size.argtypes = [c_char_p, POINTER(c_char_p), POINTER(c_longlong), POINTER(c_longlong)]
get_image_size.restype = POINTER(c_longlong)

# path = c_char_p()
ret = c_char_p()
w = c_longlong()
h = c_longlong()

path = sys.argv[1].encode('utf-8')

# err = get_image_size(path, ret, byref(w), byref(h))
get_image_size(path, byref(ret), byref(w), byref(h))
# if err != 200:
#     print("error: %s" % (err))
#     sys.exit(1)

print("%s: %dx%d" % (path, w.value, h.value))
print("%s" % (ret))
print(ret.value)
