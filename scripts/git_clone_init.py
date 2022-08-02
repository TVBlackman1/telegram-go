import shutil
from pathlib import Path
import os
import subprocess

def is_exist(path: str) -> bool:
    _path = Path(path)
    return _path.is_file()

def execute_bash(command: str):
    process = subprocess.Popen(command.split(), stdout=subprocess.PIPE)
    output, error = process.communicate()

def restore_missing_env_files():
    main_filename = './example.env'
    filenames = ('./develop.env', './test.env', './production.env')

    if not is_exist(main_filename):
        print(f'{main_filename} is not exists')
        exit()
    for filename in filenames:
        if not is_exist(filename):
            print(f'{filename} created')
            shutil.copy(main_filename, filename)

def download_dependencies():
    execute_bash("go mod download")

def place_dependencies_locally():
    execute_bash("go mod vendor")

def main():
    os.chdir('..')
    restore_missing_env_files()
    download_dependencies()
    place_dependencies_locally()

if __name__ == '__main__':
    main()