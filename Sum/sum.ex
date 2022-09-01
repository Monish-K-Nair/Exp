defmodule Sum do

  def main() do
    a = IO.gets("") |> String.trim |> String.to_integer
    b = IO.gets("") |> String.trim |> String.to_integer
    (a + b) |> IO.puts
  end
end

Sum.main()
