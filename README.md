# GURL

GURL is a go and redis based url shortener.

## Getting started

```bash
docker run 
```

## Configuration

| Parameter                    | Env                       | config.yaml                         | Default          |
|------------------------------|---------------------------|-------------------------------------|------------------|
| --port=9000                  | PORT                      | port: int                           | 9000             |
| --db.addr=""                 | DB_ADDR                   | db.addr: string                     | localhost:6379   |
| --db.database=               | DB_DATABASE               | db.database: int                    | 0                |
| --db.password=               | DB_PASSWORD               | db.password: string                 |                  |
| --page.title=                | PAGE_TITLE                | page.title: string                  | GURL - Shortener |
| --page.lang=                 | PAGE_LANG                 | page.lang: string                   | en               |
| --page.name=                 | PAGE_NAME                 | page.name: string                   | GURL             |
| --page.color                 | PAGE_COLOR                | page.color: string                  | blue             |
| --page.footer.header=        | PAGE_FOOTER_HEADER        | page.footer.header: string          |                  |
| --page.footer.copyyear=      | PAGE_FOOTER_COPYYEAR      | page.footer.copyyear: int           |                  |
| --page.footer.copystatement= | PAGE_FOOTER_COPYSTATEMENT | page.footer.copystatement: string   |                  |
| --page.footer.content=       | PAGE_FOOTER_CONTENT       | page.footer.content: string         |                  |
|                              |                           | page.footer.links: Array\<Link>     | []               |
|                              |                           | page.footer.morelinks: Array\<Link> | []               |

### Objects

#### DB

| Key      | Type   | Default        |
|----------|--------|----------------|
| Addr     | string | localhost:6379 |
| Database | int    | 0              |
| Password | string ||

#### Page

| Key   | Type   | Default          |
|-------|--------|------------------|
| Title | string | GURL - Shortener |
| Lang  | string | en               |
| Name  | string | GURL             |
| Color | string | blue             |

#### Footer

| Key           | Type         | Default |
|---------------|--------------|---------|
| Header        | string       ||
| CopyYear      | int          ||
| CopyStatement | string       ||
| Content       | string       ||
| Links         | Array\<Link> | []      |
| MoreLinks     | Array\<Link> | []      |

#### Link

| Key    | Type   | Default |
|--------|--------|---------|
| Text   | string ||
| Target | string ||
