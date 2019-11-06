Array.class_eval do
	def my_map(&block)
		mapped_array = []
    self.each { |e| mapped_array.push(block.call(e)) }
    mapped_array
	end
end


module Cat
  def self.included(base)
    base.extend(ClassMethods)
  end

  module ClassMethods
    def meow(s)
      puts "meow " + s
    end
  end
end

module A
  include Cat
end

class B < A
  meow "this"
end

module Mod
  include Cat

  meow "this"
end


# def lcs(str1, str2)
#   ## The algorithm can be described as: 
#   ## 1. Find match
#   ## 2. If found, continue to find sequences
#   ## 3. Store each sequence separately
#   ## 4. Check for sequence from the last one
#   ## 5. If a sequence does not continues, pop it off, go back to the previous one
#   ## 6. Present the largest sequence found as the end result
  
#   i, j, e = 0, 0, ""
#   r, s, t = [], [-1], []
#   c = false

#   while true
#     break if i == str1.length
#     c ? e = t[t.length-1] : e = ""
#     j = s[s.length-1] + 1
#     while true
#       break if j == str2.length
#       if str1[i] == str2[j]
#         e += str1[i]
#         c = true
#         s << j
#         r << e
#         t << e
#         break
#       end
#       j += 1
#     end
#     unless c
#       t.pop
#       s.pop if s.length > 1
#     end
#     c = false if s.length == 1 ## start over, not a continuation anymore 
#     i += 1
#   end

#   p "r: #{r}"
#   r.max_by { |x| x.length }
# end

# ABCDGH
# EADBHR