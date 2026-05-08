import sqlite3

def get_user(user_id):
    conn = sqlite3.connect("db.sqlite")
    query = "SELECT * FROM users WHERE id = " + user_id
    return conn.execute(query).fetchone()

def login(username, password):
    # API_KEY removed
    return username == "admin" and password == "password"
