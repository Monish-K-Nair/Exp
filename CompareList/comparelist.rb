a = gets.rstrip.split.map(&:to_i)
b = gets.rstrip.split.map(&:to_i)
x = y = 0
a.each_with_index do |val, index| 
    if val > b[index]
       x += 1
    end
    if val < b[index]
       y += 1
    end
end
p [x,y]