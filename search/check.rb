def bsearch(col, el, l, r, m)
  puts "l: #{l}"
  puts "r: #{r}"
  puts "m: #{m}"
  puts " ----- "
  if el == col[m]
    puts "col[m]: #{col[m]}"
    puts "m: #{m}"
    return m 
  end
	if el < col[m]
		r = m - 1
		m = (l + r) / 2
		m = bsearch(col, el, l, r, m)
  end
	if el > col[m]
		l = m + 1
		m = (l + r) / 2
		m = bsearch(col, el, l, r, m)
  end
	m
end