counter = 0

wrk.method = "POST"

request = function()
  path = "/foo" .. counter .. "/bar" .. counter
  counter = counter + 1
  return wrk.format(nil, path)
end
