#!/usr/bin/env python3
import getpass
import hashlib


class User:
    def __init__(self, email, user, passwrd_hash, name, role) -> None:
        self.email = email
        self.username = user
        self.passwrd_hash = passwrd_hash
        self.name = name
        self.role = role

    def __repr__(self) -> str:
        return "Email:    {}\nUsername: {}\nName:     {}\nRole:     {}".format(
            self.email, self.username, self.name, self.role)


if __name__ == "__main__":
    email = input("Email: ")
    username = input("Username: ")
    passwrdHash = hashlib.md5(getpass.getpass("Password: ").encode())
    name = input("Fullname: ")
    role = input("Role: ")

    usr = User(email, username, passwrdHash, name, role)
    print(usr)
    print("Password Hash:", passwrdHash)
