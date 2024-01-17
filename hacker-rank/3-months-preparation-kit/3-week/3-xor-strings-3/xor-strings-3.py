def strings_xor(s, t):
    res = ""
    for i in range(len(s)):
        if s[i] == t[i]:
            res += '0';
        else:
            res += '1';

    return res

s = input()
t = input()
print(strings_xor(s, t))

# El ejercicio en Golang no funciona pero trata de debuggear el metodo el cual hice los siguientes cambios
# Cambio 1:
#  if s[i] = t[i]:
# por 
#  if s[i] == t[i]: 
#Cambio 2:
#res = '0';
#por
#res += '0';
#Cambio 3:
#res = '1';
#por
#res += '1';

