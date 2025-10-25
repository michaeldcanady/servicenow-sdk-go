FROM docker.io/squidfunk/mkdocs-material

RUN pip install pymdown-extensions
RUN pip install markdown_include
