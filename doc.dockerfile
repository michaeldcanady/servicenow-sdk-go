FROM docker.io/squidfunk/mkdocs-material

COPY . /docs

COPY requirements.txt requirements.txt

RUN pip install --no-cache-dir -r requirements.txt

CMD ["serve", "--dev-addr=0.0.0.0:8000"]
