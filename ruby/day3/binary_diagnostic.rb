# frozen_string_literal: true

require_relative '../aoc'

inputs = Aoc::Input.new(3).lines.map { _1.to_i(2) }
counts = []

inputs.each do |input|
  current_bit = 0

  while input.positive?
    counts[current_bit] ||= 0
    counts[current_bit] += 1 if input.odd?
    current_bit += 1
    input = input >> 1
  end
end

gamma = counts.map { _1 > inputs.size / 2 ? 1 : 0 }.reverse.join.to_i(2)
epsilon = counts.map { _1 < inputs.size / 2 ? 1 : 0 }.reverse.join.to_i(2)

puts "The first answer is #{gamma * epsilon}"
