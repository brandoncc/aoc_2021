inputs_path = File.expand_path("#{File.dirname(__FILE__)}../../inputs.txt")
lines = File.read(inputs_path).lines.map(&:to_i)

last = lines.first
count = 0

lines.each do |l|
  count += 1 if l > last
  last = l
end

puts "The answer is #{count}"
