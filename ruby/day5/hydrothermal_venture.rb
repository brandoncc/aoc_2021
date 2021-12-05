# frozen_string_literal: true

require_relative '../aoc'

def find_largest_x(inputs)
  inputs.flat_map { [_1[:x1], _1[:x2]] }.max
end

def find_largest_y(inputs)
  inputs.flat_map { [_1[:y1], _1[:y2]] }.max
end

def velocity(value1, value2)
  return 0 if value1 == value2
  return -1 if value1 > value2

  1
end

def covered_points(input)
  x_velocity = velocity(input[:x1], input[:x2])
  y_velocity = velocity(input[:y1], input[:y2])

  x = input[:x1] - x_velocity
  y = input[:y1] - y_velocity

  points = []

  loop do
    x += x_velocity
    y += y_velocity
    points << [x, y]

    break if x == input[:x2] && y == input[:y2]
  end

  points
end

def horizontal?(input)
  input[:y1] == input[:y2]
end

def covered_points_vertical(input)
  small, large = [input[:y1], input[:y2]].sort
  distance = large - small + 1

  distance.times.map do |y_offset|
    [input[:x1], small + y_offset]
  end
end

def covered_points_horizontal(input)
  small, large = [input[:x1], input[:x2]].sort
  distance = large - small + 1

  distance.times.map do |x_offset|
    [small + x_offset, input[:y1]]
  end
end

def print_table(locations)
  puts(locations.map do |row|
    row.map { |v| v.zero? ? '.' : v }.join('')
  end)
end

def straight_only(inputs)
  inputs.select { _1[:x1] == _1[:x2] || _1[:y1] == _1[:y2] }
end

inputs = Aoc::Input.new(5).lines

inputs = inputs.join(' ').scan(/\d+/).map(&:to_i).each_slice(4).map do |point|
  { x1: point[0], y1: point[1], x2: point[2], y2: point[3] }
end

def overlap_count(inputs)
  locations = []

  largest_x = find_largest_x(inputs)
  largest_y = find_largest_y(inputs)

  (largest_y + 1).times { locations << Array.new(largest_x + 1, 0) }

  inputs.each do |input|
    covered_points(input).each do |point|
      row = locations[point[1]]
      row[point[0]] += 1
    end
  end

  locations.flatten.count { _1 > 1 }
end

puts "The first solution is #{overlap_count(straight_only(inputs))}"
puts "The second solution is #{overlap_count(inputs)}"
