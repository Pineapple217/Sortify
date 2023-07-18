from datetime import datetime
import time
import pandas as pd
import spotipy
from spotipy.oauth2 import SpotifyOAuth
from dotenv import load_dotenv
import os
import csv
import typer
from typing_extensions import Annotated
import random
from rich import print
from rich.console import Console
from rich.progress import Progress, SpinnerColumn, TextColumn

from ml import birch_cluster

app = typer.Typer(no_args_is_help=True)
err_console = Console(stderr=True)


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
            scope=[
                "playlist-read-private",
                "user-library-read",
                "playlist-modify-public",
            ],
        )
    )
    return sp


@app.command()
def export_liked():
    """
    Exports liked songs from Spotify to a csv file.
    """
    current_datetime = datetime.now()
    formatted_datetime = current_datetime.strftime("%Y-%m-%d_%H-%M-%S")
    if not os.path.exists("csv"):
        os.makedirs("csv")
    filename = f"./csv/liked_{formatted_datetime}.csv"
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

    def process_result(results):
        ids = [t["track"]["id"] for t in results["items"]]
        featueres = sp.audio_features(ids)

        for f, track_id in zip(featueres, ids):
            try:
                writer.writerow(dict((k, f[k]) for k in fieldnames))
            except:
                d = dict((k, None) for k in fieldnames)
                d["id"] = track_id
                writer.writerow(d)

    results = sp.current_user_saved_tracks(limit=50)  # 50 is max
    with Progress() as progress:
        task1 = progress.add_task("Pulling liked songs...", total=results["total"])
        progress.console.print(
            f"Found {results['total']} liked song{'' if results['total'] == 1 else 's'}"
        )

        progress.update(task1, completed=results["offset"] + results["limit"])
        process_result(results)
        while results["next"]:
            results = sp.next(results)
            process_result(results)
            progress.update(task1, completed=results["offset"] + results["limit"])


@app.command()
def import_playlist(
    playlist_csv,
    random_order: Annotated[
        bool,
        typer.Option(
            "--random", "-R", help="Shuffles the playlist into a random order."
        ),
    ] = False,
):
    """
    Import a playlist from a csv file.
    """
    if not os.path.exists(playlist_csv):
        print(f"file {playlist_csv} not found.")
        raise typer.Exit(code=1)
    user_id = sp.me()["id"]
    playlist_id = sp.user_playlist_create(user_id, "AI_TEST1")
    max_length = 100  # 100 api limit for adding tracks
    try:
        id_list = pd.read_csv(playlist_csv)["id"]
    except:
        err_console.print(f"file {playlist_csv} does not contain an id-column.")
        raise typer.Exit(code=1)
    if random_order:
        random.shuffle(id_list)

    lists = [id_list[i : i + max_length] for i in range(0, len(id_list), max_length)]
    for ids in lists:
        sp.user_playlist_add_tracks(user_id, playlist_id["id"], ids)


@app.command()
def brich_playlist_gen(
    liked_csv: Annotated[str, typer.Argument(resolve_path=True)] = "Most recent csv",
    playlist_count: Annotated[
        int, typer.Option("-n", help="Amount of playlists to export.")
    ] = 15,
    random_order: Annotated[
        bool,
        typer.Option(
            "--random", "-R", help="Shuffles the playlists songs into a random order."
        ),
    ] = False,
):
    """
    Generates playlists using the BIRCH algorithm.
    """
    if liked_csv == "Most recent csv":
        folder_path = "./csv"
        csv_files = [
            f
            for f in os.listdir(folder_path)
            if f.startswith("liked_") and f.endswith(".csv")
        ]
        if len(csv_files) == 0:
            err_console.print(
                "[bold red]Alert![/bold red] No liked csv found, use the `export-liked` command to generate this file."
            )
            raise typer.Exit(code=1)
        csv_files.sort(
            key=lambda x: os.path.getmtime(os.path.join(folder_path, x)), reverse=True
        )
        liked_csv = os.path.join(folder_path, csv_files[0])
    # print(f"Using [italic white]{liked_csv}[/italic white] for playlist generation")

    with Progress(
        SpinnerColumn(),
        TextColumn("[progress.description]{task.description}"),
        transient=True,
    ) as progress:
        task_clustering = progress.add_task(description="Clustering...", total=1)
        df = pd.read_csv(liked_csv, index_col="id")
        playlists = birch_cluster(df, playlist_count)
        progress.update(task_clustering, completed=1)

        task_playlist = progress.add_task(description="Creating playlists...", total=1)
        playlists.reset_index(inplace=True)
        playlists.set_index("playlist", inplace=True)

        user_id = sp.me()["id"]

        max_length = 100  # 100 api limit for adding tracks
        for playlist_nr, playlist_df in playlists.groupby(level=0):
            playlist = sp.user_playlist_create(
                user=user_id,
                name=f"AI_{playlist_nr}",
                description="Generated using the Birch Algorithm",
            )
            id_list = playlist_df["id"].to_list()
            if random_order:
                random.shuffle(id_list)
            lists = [
                id_list[i : i + max_length] for i in range(0, len(id_list), max_length)
            ]
            for ids in lists:
                sp.user_playlist_add_tracks(user_id, playlist["id"], ids)

        progress.update(task_playlist, completed=1)


sp = get_sp_client()

if __name__ == "__main__":
    app()
