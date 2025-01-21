import os
import pathlib
from setuptools import setup
from setuptools import find_packages


def read(rel_path):
    here = os.path.abspath(os.path.dirname(__file__))
    with open(os.path.join(here, rel_path)) as fp:
        return fp.read()


def get_version(rel_path):
    for line in read(rel_path).splitlines():
        if line.startswith("__version__"):
            delim = '"' if '"' in line else "'"
            return line.split(delim)[1]
    raise RuntimeError("Unable to find version string.")


# Read the requirements from the requirements.txt file
with open('requirements.txt') as f:
    required = f.read().splitlines()

HERE = pathlib.Path(__file__).parent
README = (HERE / "README.md").read_text()
VERSION = get_version("eai/version.py")

setup(
    name="eai",
    version=VERSION,
    author="EternalAI",
    description="Toolkit to deploy model onchain",
    long_description=README,
    long_description_content_type="text/markdown",
    license='LICENSE.txt',
    packages=find_packages(
        include=("eai", "eai.*"),
    ),
    install_requires=required,
    classifiers=['Operating System :: POSIX', ],
    entry_points={
        'console_scripts': [
            'eai = eai.cli:main',
        ]
    },
    python_requires='>=3.10',
)
