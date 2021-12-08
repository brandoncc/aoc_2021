# frozen_string_literal: true

require_relative '../aoc'

inputs = Aoc::Input.new(7).lines.first.split(',').map(&:to_i).sort

def solution1(inputs)
  min = inputs.min
  max = inputs.max

  best = nil

  movements = []

  inputs.each_with_index do |inp, ind1|
    movements[ind1] = (min..max).map do |v|
      (v - inp).abs
    end
  end

  movements[0].count.times do |col|
    sum = movements.map { |m| m[col] }.inject(:+)
    best = sum if best.nil? || best > sum
  end

  best + min
end

def solution2(inputs)
  min = inputs.min
  max = inputs.max

  best = nil

  movements = []

  inputs.each_with_index do |inp, ind1|
    movements[ind1] = (min..max).map do |v|
      distance = (v - inp).abs
      distance * (1 + distance) / 2
    end
  end

  movements[0].count.times do |col|
    sum = movements.map { |m| m[col] }.inject(:+)
    best = sum if best.nil? || best > sum
  end

  best + min
end

puts "The first answer is #{solution1(inputs)}"
puts "The first answer is #{solution2(inputs)}"
