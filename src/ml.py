import pandas as pd

from sklearn.compose import make_column_selector, make_column_transformer
from sklearn.impute import SimpleImputer
from sklearn.pipeline import FunctionTransformer, make_pipeline
from sklearn.preprocessing import OneHotEncoder, StandardScaler

from sklearn.cluster import Birch
from sklearn.metrics import silhouette_score

num_pipeline = make_pipeline(
    StandardScaler(),
)

cat_pipeline = make_pipeline(
    OneHotEncoder(handle_unknown="ignore"),
)

pre_pros = make_column_transformer(
    (
        num_pipeline,
        [
            "acousticness",
            "danceability",
            "speechiness",
            "instrumentalness",
            "valence",
            "energy",
            "liveness",
            "tempo",
        ],
    ),
    (cat_pipeline, ["time_signature", "key", "mode"]),
    remainder="passthrough",
)


def birch_cluster(df: pd.DataFrame, n_clusters: int) -> pd.DataFrame:
    df_birch = df.copy()
    df_birch.dropna(inplace=True)
    df_birch.drop(["duration_ms", "loudness"], inplace=True, axis=1)
    brc = Birch(threshold=0.1, n_clusters=n_clusters)

    y_pred = make_pipeline(pre_pros, brc).fit_predict(df_birch)
    # score = silhouette_score(df_birch, brc.labels_)

    df_birch["playlist"] = y_pred
    return df_birch
