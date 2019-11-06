##
## Find the substring in a string.
##
# def print_combinations(str)
#   r = [str]
#   str_len = str.length
#   return r if str_len <= 1
#   i, j = 0, 0
#   str_fact = fact(str_len)
#   circular = false
#   s = ""
  
#   while true
#     break if r.length == str_fact
#     if i == str_len - 1
#       ## time for circular moves
#       puts "circular"
#       unless circular
#         j = s.length - 1
#         s = str.clone
#       end
#       s = s[j] + s[0...s.length - 1]
#       break if s == str
#       r << s
#       circular = true
#     else
#       s = str.clone
#       puts "i: #{i}, j: #{j}"
#       if i < j
#         puts "str[i]: #{str[i]}, str[j]: #{str[j]}"
#         a = str[i]
#         b = str[j]
#         puts "s[i]: #{s[i]}, s[j]: #{s[j]}"
#         s[i] = b
#         s[j] = a
#         r << s
#       end
#       if j == str_len - 1
#         i += 1
#         i <= str_len - 1 ? j = i : j = i + 1
#       else
#         j += 1
#       end
#       puts "s: #{s}"
#       puts " ---------- "
#       s = "" if i == str_len - 1
#     end
#   end

#   puts r.length
#   p r
#   puts r.uniq.length
#   p r.uniq
# end

## Incomplete
def find_combinations(str)
  ## Find all combinations of characters of a string
  r = [str]
  str_len = str.length
  return r if str_len <= 1
  i, j, k, l = 0, 0, 0, 0
  s = ""
  imd = []
  str_fact = fact(str_len)

  while true
    break if (r + imd).length == str_fact
    t = []
    circular_init = true
    j += 1 if j == i
    k < j ? k = j + 1 : k += 1
    s = r[l].clone
    
    while true
      if j == str_len - 1
        ## time for circular moves
        puts "circular"
        if circular_init
          k = str_len - 1
          s = r[l].clone
        end
        s = s[k] + s[0...s.length - 1]
        break if s == r[l]
        t << s
        circular_init = false
      else
        puts "i: #{i}, j: #{j}, k: #{k}"
        if i != j && j != k && k != i
          puts "str[j]: #{str[j]}, str[k]: #{str[k]}"
          a = str[j]
          b = str[k]
          puts "s[j]: #{s[j]}, s[k]: #{s[k]}"
          s[j] = b
          s[k] = a
          t << s
        end
        if k == str_len - 1
          j += 1
          k = j + 1 if j < str_len - 1
        else
          k += 1
        end
        puts "s: #{s}"
        puts " ---------- "
        s = "" if j == str_len - 1
      end
    end

    unless i == str_len - 1
      s1 = r[l].clone
      x = i
      y = x + 1
      a = r[l][x]
      b = r[l][y]
      s1[x] = b
      s1[y] = a
      r << s1
      i = y
      l += 1
      j, k = 0, 0
    end

    puts t.length
    p t
    puts ""
    imd << t
    imd.flatten!
    puts " ============ "
  end

  r += imd
  puts r.length
  p r
  puts ""
  puts r.uniq.length
  p r.uniq
end

def fact(n)
  m = 1
  n.downto(1) { |i| m *= i }
  m
end

# ABC ACB CBA BAC CAB BCA 

## For ABCD (4! = 24):
#  With Interchanges  |  Circular
#  -----------------  |  --------
# ABCD ACBD ADCB ABDC | ADBC ACDB
# BACD CABD DACB BADC | BDAC BCDA
# CBAD BCAD DBAC CDAB | DCBA ADCB
# CBDA BCDA DBCA CDBA | ACBD DACB  