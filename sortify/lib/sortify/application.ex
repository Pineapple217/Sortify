defmodule Sortify.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      # Start the Telemetry supervisor
      SortifyWeb.Telemetry,
      # Start the Ecto repository
      Sortify.Repo,
      # Start the PubSub system
      {Phoenix.PubSub, name: Sortify.PubSub},
      # Start Finch
      {Finch, name: Sortify.Finch},
      # Start the Endpoint (http/https)
      SortifyWeb.Endpoint
      # Start a worker by calling: Sortify.Worker.start_link(arg)
      # {Sortify.Worker, arg}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Sortify.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    SortifyWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
