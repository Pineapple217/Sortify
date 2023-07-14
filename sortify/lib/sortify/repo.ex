defmodule Sortify.Repo do
  use Ecto.Repo,
    otp_app: :sortify,
    adapter: Ecto.Adapters.Postgres
end
