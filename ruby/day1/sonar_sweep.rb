# frozen_string_literal: true

require_relative '../aoc'

class NoisySolution
  def solution
    inputs = Aoc::Input.new(1).lines.map(&:to_i)
    last = inputs.first
    count = 0

    inputs.each do |l|
      count += 1 if l > last
      last = l
    end

    count
  end
end

class WindowedSolution
  def solution
    inputs = Aoc::Input.new(1).lines.map(&:to_i)
    last = inputs.first(3).inject(:+)
    count = 0

    (3...inputs.size).each do |i|
      next_val = inputs[i] + inputs[i - 1] + inputs[i - 2]
      count += 1 if next_val > last
      last = next_val
    end

    count
  end
end

puts "The first answer is #{NoisySolution.new.solution}"
puts "The second answer is #{WindowedSolution.new.solution}"
