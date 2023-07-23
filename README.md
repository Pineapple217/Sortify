# Sortify

AI driver playlist maker of liked songs

## CLI Installation

### Requirements

python 3.11+

### Instructions

1. Clone the repo

```sh
git clone https://github.com/Pineapple217/Sortify.git
```

2. Create python venv

```sh
python -m venv venv
.\venv\Scripts\activate
```

3. Instal requirements

```sh
pip install -r .\requirements.txt
```

4. Spotify API

go to https://developer.spotify.com/ and create an application.

5. Configure environment variables

Copy the `template.env` file in file in the fields

```sh
cp template.env .env
```

I recommand setting the redirectt url to `http://localhost:8888`, make sure this is the same in the Spotify dahsboard.

6. Use the aplication

```sh
python .\src\cli.py
```
