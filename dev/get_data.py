from datetime import datetime
import spotipy
from spotipy.oauth2 import SpotifyOAuth
from dotenv import load_dotenv
import os
from pprint import pprint
import csv


load_dotenv()

SPOTIFY_CLIENT_ID = os.getenv("SPOTIFY_CLIENT_ID")
SPOTIFY_CLIENT_SECRET = os.getenv("SPOTIFY_CLIENT_SECRET")
SPOTIFY_REDIRECT_URI = os.getenv("SPOTIFY_REDIRECT_URI")

sp = spotipy.Spotify(
    auth_manager=SpotifyOAuth(
        client_id=SPOTIFY_CLIENT_ID,
        client_secret=SPOTIFY_CLIENT_SECRET,
        redirect_uri=SPOTIFY_REDIRECT_URI,
        scope="user-library-read",
    )
)

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
