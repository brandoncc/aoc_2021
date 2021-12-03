# frozen_string_literal: true

require_relative '../aoc'

class Submarine
  def initialize
    @horizontal = 0
    @depth = 0
    @aim = 0
  end

  def answer
    @horizontal * @depth
  end
end

class BasicSubmarine < Submarine
  def up(amount)
    @depth -= amount
  end

  def down(amount)
    @depth += amount
  end

  def forward(amount)
    @horizontal += amount
  end
end

class AimableSubmarine < Submarine
  def up(amount)
    @aim -= amount
  end

  def down(amount)
    @aim += amount
  end

  def forward(amount)
    @horizontal += amount
    @depth += @aim * amount
  end
end

class CommandParser
  def self.parse(command)
    movement, amount = command.split

    {
      movement: movement,
      amount: amount.to_i
    }
  end
end

inputs = Aoc::Input.new(2).lines

sub1 = BasicSubmarine.new
sub2 = AimableSubmarine.new

inputs.each do |input|
  parsed = CommandParser.parse(input)

  sub1.send(*parsed.values)
  sub2.send(*parsed.values)
end

puts "The first answer is #{sub1.answer}"
puts "The second answer is #{sub2.answer}"
