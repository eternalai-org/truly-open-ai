# Use the official Python 3.10 image
FROM python:3.10

# Set the working directory in the container
WORKDIR /app

# Copy the Pipfile and Pipfile.lock to the working directory
COPY Pipfile Pipfile.lock /app/

# Install Pipenv and project dependencies
RUN pip install pipenv && pipenv install --deploy --ignore-pipfile

# Copy the source code to the working directory
COPY src /app/src

# Expose port 5000
EXPOSE 5000

# Set environment variables
ENV FLASK_APP=src/main.py
ENV FLASK_RUN_HOST=0.0.0.0

# Install Gunicorn
RUN pipenv install gunicorn

# Command to run the Flask application with Gunicorn
CMD ["pipenv", "run", "gunicorn", "--chdir", "/app/src", "--bind", "0.0.0.0:5000", "src.main:app"]
