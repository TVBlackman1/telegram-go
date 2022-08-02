# pip3 install dotenv

from importlib.resources import path
import shutil
import os
import subprocess
from dotenv import dotenv_values

def get_config_content(db: str, port: int, user: str, password: str):
    return f"""
flyway.url=jdbc:postgresql://postgres:{port}/{db}
flyway.user={user}
flyway.password={password}
flyway.baselineOnMigrate=false
"""

def get_filepath():
    path_to_file = os.path.join('..', 'configs', 'flyway.config')
    return os.path.abspath(path_to_file)

def main():
    envs = dotenv_values("../example.env")
    db = envs["POSTGRES_DBNAME"]
    port = envs["POSTGRES_PORT"]
    user = envs["POSTGRES_USER"]
    password = envs["POSTGRES_PASS"]

    text = get_config_content(db, port, user, password)
    path_to_file = get_filepath()
    with open(path_to_file, 'w') as file:
        file.write(text)
if __name__ == '__main__':
    main()