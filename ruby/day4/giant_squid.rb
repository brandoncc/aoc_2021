# frozen_string_literal: true

require_relative '../aoc'

class Board
  def initialize(rows)
    setup_board(rows)
    @drawn = []
  end

  def to_s
    @rows.inspect
  end

  def call_number(number)
    @drawn << number
  end

  def won?
    winning_location
  end

  def points
    won? ? calculate_points : 0
  end

  def winning_call
    @drawn.last.to_i
  end

  private

  def setup_board(rows)
    @rows = rows.map { _1.strip.split(/ +/) }
  end

  def columns
    @columns ||= @rows.first.size.times.map do |i|
      @rows.map { _1[i] }
    end
  end

  def winning_row
    @rows.find { |r| r.uniq.count == (r & @drawn).count }
  end

  def winning_column
    columns.find { |c| c.uniq.count == (c & @drawn).count }
  end

  def winning_location
    winning_row || winning_column
  end

  def board_points
    @rows
      .flatten
      .reject { |e| @drawn.include?(e) }
      .map(&:to_i)
      .reduce(:+)
  end

  def winning_location_points
    winning_location&.map(&:to_i)&.inject(:+)
  end

  def calculate_points
    board_points
  end
end

inputs = Aoc::Input.new(4).lines
draws = inputs.shift(2).first.split(',')
boards = []

while inputs.any?
  boards << Board.new(inputs.shift(5))
  inputs.shift
end

winners = []

draws.each do |d|
  break if winners.count == boards.count

  boards.each do |b|
    next if winners.include?(b)

    b.call_number(d)
    winners << b if b.won?
  end
end

puts "The first solution is #{winners.first.points * winners.first.winning_call}"
puts "The second solution is #{winners.last.points * winners.last.winning_call}"
