# Copyright 2023 parkminwoo, MIT License
from setuptools import find_packages
from setuptools import setup

version = {}

with open("dagent/version.py") as fp:
    exec(fp.read(), version)

def get_long_description():
    with open("README.md", encoding="UTF-8") as f:
        return f.read()

dependencies = [
    "schedule==1.2.2",
    "python-dotenv==1.0.1",
    "requests==2.32.3",
    "pydantic==2.9.2",
    "bs4==0.0.2",
    "singleton-decorator==1.0.0"
]

setup(
    name="dagent",
    version=version["__version__"],
    author="EternalAI",
    author_email="dev@eternalai.org",
    description="EternalAI Agent is a conversational decentralized agent that can perform various tasks using a combination of toolsets and language models. The agent can be configured to execute multiple missions, each with its own set of toolsets and language models.",
    long_description=get_long_description(),
    long_description_content_type="text/markdown",
    url="https://github.com/eternalai-org/dagent",
    packages=find_packages(exclude=["dagent.service.py"], include=["dagent", "dagent.*"]),
    python_requires=">=3.10.0",
    install_requires=dependencies,
    keywords="Python, LLM, Decentralized AI, Modular design",
    classifiers=[
        "Development Status :: 5 - Production/Stable",
        "Intended Audience :: Science/Research",
        "Natural Language :: English",
        "Programming Language :: Python",
        "Programming Language :: Python :: 3.10",
        "License :: OSI Approved :: MIT License",
        "Topic :: Scientific/Engineering :: Artificial Intelligence",
    ],
    license="MIT",
)