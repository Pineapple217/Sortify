from datetime import datetime
import pandas as pd
import spotipy
from spotipy.oauth2 import SpotifyOAuth
from dotenv import load_dotenv
import os
from pprint import pprint
import csv
import typer

app = typer.Typer()


def get_sp_client():
    load_dotenv()

    SPOTIFY_CLIENT_ID = os.getenv("SPOTIFY_CLIENT_ID")
    SPOTIFY_CLIENT_SECRET = os.getenv("SPOTIFY_CLIENT_SECRET")
    SPOTIFY_REDIRECT_URI = os.getenv("SPOTIFY_REDIRECT_URI")

    sp = spotipy.Spotify(
        auth_manager=SpotifyOAuth(
            client_id=SPOTIFY_CLIENT_ID,
            client_secret=SPOTIFY_CLIENT_SECRET,
            redirect_uri=SPOTIFY_REDIRECT_URI,
            scope="playlist-modify-public",
        )
    )
    return sp


@app.command()
def export_liked():

    current_datetime = datetime.now()
    formatted_datetime = current_datetime.strftime("%Y-%m-%d_%H-%M-%S")
    if not os.path.exists("csv"):
        os.makedirs("csv")
    filename = f"./csv/{formatted_datetime}.csv"
    file = open(filename, mode="w", newline="")
    fieldnames = [
        "id",
        "acousticness",
        "danceability",
        "speechiness",
        "key",
        "loudness",
        "instrumentalness",
        "valence",
        "duration_ms",
        "mode",
        "energy",
        "liveness",
        "tempo",
        "time_signature",
    ]
    writer = csv.DictWriter(file, fieldnames=fieldnames)

    writer.writeheader()

    def show_tracks(results):
        ids = [t["track"]["id"] for t in results["items"]]
        featueres = sp.audio_features(ids)

        for f, track_id in zip(featueres, ids):
            try:
                writer.writerow(dict((k, f[k]) for k in fieldnames))
            except:
                d = dict((k, None) for k in fieldnames)
                d["id"] = track_id
                writer.writerow(d)
        pprint(ids)

    results = sp.current_user_saved_tracks(limit=50)  # 50 is max
    show_tracks(results)

    while results["next"]:
        results = sp.next(results)
        show_tracks(results)


@app.command()
def import_playlist(playlist_csv):
    user_id = sp.me()["id"]
    playlist_id = sp.user_playlist_create(user_id, "AI_TEST1")
    max_length = 100
    id_list = pd.read_csv(playlist_csv)["id"]

    lists = [id_list[i : i + max_length] for i in range(0, len(id_list), max_length)]
    for ids in lists:
        sp.user_playlist_add_tracks(user_id, playlist_id["id"], ids)


sp = get_sp_client()

if __name__ == "__main__":
    app()
