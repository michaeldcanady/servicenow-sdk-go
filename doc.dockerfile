FROM docker.io/squidfunk/mkdocs-material

COPY requirements.txt requirements.txt

RUN pip install --no-cache-dir -r requirements.txt

CMD ["serve", "--dirtyreload", "--dev-addr=0.0.0.0:8000"]
