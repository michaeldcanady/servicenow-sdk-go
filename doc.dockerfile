FROM docker.io/squidfunk/mkdocs-material

RUN pip install pymdown-extensions
RUN pip install markdown_include

CMD ["serve", "--dirtyreload", "--dev-addr=0.0.0.0:8000"]
