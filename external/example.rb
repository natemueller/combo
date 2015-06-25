#!/usr/bin/env ruby

# imeplements a random  player conforming to the external player interface

require "json"

# calculate our list of possible moves
def available_moves(cmd)
  moves = []

  cmd['board']['squares'].flatten.each do |sq|
    if sq['piece_color'] != cmd['color'] || sq['piece_count'] == 0
      next
    end

    (-1..1).zip(-1..1) do |dx, dy|
      if dx == 0 && dy == 0
        next
      end

      x, y = sq["x"], sq["y"]

      1.upto(sq["piece_count"]) do |distance|
        x += dx
        y += dy

        if x < 0 || x >= cmd['board']['width'] || y < 0 || y >= cmd['board']['height']
          break
        end

        # "p"otential square
        psq = cmd['board']['squares'][x][y]

        distance.upto(sq['piece_count']) do |split_size|
          if split_size > 1 || psq['piece_count'] == 0 || psq['piece_color'] == cmd['color']
            moves << {
              "from" => {"x" => sq["x"], "y" => sq["y"]},
              "to" => {"x" => psq["x"], "y" => psq["y"]},
              "piece_count" => split_size,
            }
          end
        end

        if psq["piece_count"] > 0
          break
        end
      end
    end
  end

  return moves
end

STDIN.each_line do |line|
  moves = available_moves(JSON[line])
  STDOUT.puts(JSON[moves[rand(moves.length)]])
  STDOUT.flush
end