# frozen_string_literal: true

# Utilities related to Advent of Code
module Aoc
  # A helper class to read input files
  class Input
    attr_reader :lines

    def initialize(day)
      @day = day
      @lines = File.read(path).lines
    end

    private

    def validate_path
      raise "File not found: #{path}" unless File.exist?(path)
    end

    def path
      File.expand_path("#{File.dirname(__FILE__)}/../inputs/day#{@day}.txt")
    end
  end
end
