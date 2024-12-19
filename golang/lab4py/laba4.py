print("Daria Hljupina")
import math

def calculate(x, a):
    y = math.tan(math.pow(math.log10(a + x), 3)) / math.pow((a + x), 2.0 / 7.0)
    return y

def calculate_b(x_values, a):
    spb = []
    for x in x_values:
        spb.append(calculate(x, a))
    return spb

def calculate_a(xn, xk, delta_x, a):
    spa = []
    x = xn
    while x <= xk:
        spa.append(calculate(x, a))
        x += delta_x
    return spa

def ready_labaa4ab():
    print("A")
    print(calculate_a(1.08, 1.88, 0.16, 2))
    print("B")
    x_values = [1.16, 1.35, 1.48, 1.52, 1.96]
    print(calculate_b(x_values, 2.0))

ready_labaa4ab()