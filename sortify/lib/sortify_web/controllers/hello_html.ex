defmodule SortifyWeb.HelloHTML do
  use SortifyWeb, :html

  # def index(assigns) do
  #   ~H"""
  #   Hello :3
  #   """
  # end

  embed_templates "hello_html/*"

end
