#!/bin/bash
import shutil
from pathlib import Path
import os

def is_exist(path: str) -> bool:
    _path = Path(path)
    return _path.is_file()

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

os.chdir('..')
restore_missing_env_files()