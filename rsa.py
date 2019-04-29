#!/usr/bin/env sage

import random
import sage.all as sage

CHECKS_NUM = 50
KEY_LENGHT = 5

# x^y mod n
def power(x, y, n):
    result = 1
    if y == 0:
        return result
    q = x%n
    for i in range(y.bit_length()):
        if (1<<i)&y:
            result = (result*q)%n
        q = (q*q)%n
    return result


def is_prime(n):
    for i in range(CHECKS_NUM):
        a = random.getrandbits(n.bit_length())%n
        if power(a, n-1, n) != 1:
            return False
    return True


def gen_prime(n):
    while 1:
        p = random.getrandbits(n)
        if p == 0:
            continue
        if is_prime(p) and p != 1:
            return p

class RSA:
    def __init__(self):
        self.p = gen_prime(KEY_LENGHT)
        self.q = gen_prime(KEY_LENGHT)
        self.n = self.p * self.q
        self.phi = long(sage.euler_phi(self.n))
        self.e = self.__generate_e()
        self.d = self.__generate_d()

    def __generate_e(self):
        while 1:
            e = random.getrandbits(self.phi.bit_length())%self.phi
            if e == 0 or e == 1:
                continue
            if sage.gcd(e, self.phi) == 1:
                return long(e)

    def __generate_d(self):
        _, _, d = sage.xgcd(self.phi, self.e)
        if d < 0:
            return self.phi + long(d)
        return long(d)

    def encrypt(self, num):
        return power(num, self.e, self.n)

    def decrypt(self, num):
        return power(num, self.d, self.n)


if __name__=="__main__":
    while 1:
        rsa_system = RSA()
        num = random.getrandbits(3)
        encrypted_num = rsa_system.encrypt(num)
        decrypted_num = rsa_system.decrypt(encrypted_num)
        if num != decrypted_num:
            print("p = {}, is prime: {}".format(rsa_system.p), rsa_system.p in sage.Primes())
            print("q = {}, is prime: {}".format(rsa_system.q), rsa_system.q in sage.Primes())
            print("phi = {}".format(rsa_system.phi))
            print("e = ".format(rsa_system.e))
            print("d = ".format(rsa_system.d))
            print("({}**{})%{} = {}".format(num, rsa_system.e, rsa_system.n, encrypted_num))
            print("({}**{})%{} = {}".format(encrypted_num, rsa_system.d, rsa_system.n, decrypted_num))
            break
        print("OK")
