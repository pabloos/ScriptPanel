

from sys import argv

def suma(num1, num2):
    return int(num1) + int(num2)

def resta(num1, num2):
    return int(num1) - int(num2)

def mult(num1, num2):
    return int(num1) * int(num2)

def div(num1, num2):
    return int(num1) / int(num2)

opHash = {
    "sumar" : suma,
    "restar" : resta,
    "multiplicar" : mult,
    "dividir": div,
}

print(opHash[argv[1]](argv[2], argv[3]))