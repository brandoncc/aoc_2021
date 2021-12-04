# frozen_string_literal: true

require_relative '../aoc'

def one_counts(inputs)
  [].tap do |counts|
    inputs.each do |input|
      current_bit = 0

      while input.positive?
        counts[current_bit] ||= 0
        counts[current_bit] += 1 if input.odd?
        current_bit += 1
        input = input >> 1
      end
    end
  end
end

binary_inputs = Aoc::Input.new(3).lines.map(&:chomp)
inputs = binary_inputs.map { _1.to_i(2) }
counts = one_counts(inputs)

gamma = counts.map { _1 > inputs.size / 2 ? 1 : 0 }.reverse.join.to_i(2)
epsilon = counts.map { _1 < inputs.size / 2 ? 1 : 0 }.reverse.join.to_i(2)

puts "The first answer is #{gamma * epsilon}"

# To find oxygen generator rating, determine the most common value (0 or 1) in
# the current bit position, and keep only numbers with that bit in that
# position. If 0 and 1 are equally common, keep values with a 1 in the position
# being considered.

def find_most_common_bit(inputs, position, equal_value)
  ones = inputs.map { |input| input[position] }.count(1)
  zeroes = inputs.size - ones

  return equal_value if ones == zeroes

  ones > zeroes ? 1 : 0
end

def find_least_common_bit(inputs, position, equal_value)
  ones = inputs.map { |input| input[position] }.count(1)
  zeroes = inputs.size - ones

  return equal_value if ones == zeroes

  ones > zeroes ? 0 : 1
end

def oxygen_rating(inputs, current_bit)
  return inputs.first if inputs.size < 2
  return if current_bit.negative?

  common_bit = find_most_common_bit(inputs, current_bit, 1)
  puts "common_bit: #{common_bit}"

  inputs = inputs.select do |input|
    input[current_bit] == common_bit
  end

  oxygen_rating(inputs, current_bit - 1)
end

oxy_rating = oxygen_rating(inputs, binary_inputs.first.size - 1)

def co2_scrubber(inputs, current_bit)
  return inputs.first if inputs.size < 2
  return if current_bit.negative?

  common_bit = find_most_common_bit(inputs, current_bit, 1)

  inputs = inputs.reject do |input|
    input[current_bit] == common_bit
  end

  co2_scrubber(inputs, current_bit - 1)
end

co2_rating = co2_scrubber(inputs, binary_inputs.first.size - 1)

puts "The second answer is #{co2_rating * oxy_rating}"
