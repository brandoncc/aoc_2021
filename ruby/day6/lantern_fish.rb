# frozen_string_literal: true

require_relative '../aoc'

def parse_inputs
  inputs = Aoc::Input.new(6).lines.first.split(',').map(&:to_i)

  Array.new(9, 0).tap do |counts|
    inputs.each { |input| counts[input] += 1 }
  end
end

initial_counts = parse_inputs

def pass_days(starting_fish, count)
  new_counts = starting_fish.dup

  count.times do
    old_counts = new_counts.dup

    new_counts[6] = old_counts[0] + old_counts[7]
    new_counts[8] = old_counts[0]

    new_counts[0] = old_counts[1]
    new_counts[1] = old_counts[2]
    new_counts[2] = old_counts[3]
    new_counts[3] = old_counts[4]
    new_counts[4] = old_counts[5]
    new_counts[5] = old_counts[6]
    new_counts[7] = old_counts[8]
  end

  new_counts
end

def count_fish(fish)
  fish.inject(:+)
end

solution1 = pass_days(initial_counts, 80)
puts "fish counts: #{count_fish(solution1)}"

solution2 = pass_days(initial_counts, 256)
puts "fish counts: #{count_fish(solution2)}"
