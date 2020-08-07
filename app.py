# -*- coding: utf-8 -*-
import taos
import time


def run_tdengine():
    conn = taos.connect(host="tdengine", user="root", password="taosdata", config="/etc/taos")

    # Generate a cursor object to run SQL commands
    c1 = conn.cursor()

    # Create a database named db
    try:
        c1.execute('create database pytest;')
        print('create database pytest ok!')
    except Exception as err:
        conn.close()
        raise (err)


def run_echo():
    while True:
        print('hello world!')
        time.sleep(3)


if __name__ == '__main__':
#     run_echo()
    run_tdengine()
