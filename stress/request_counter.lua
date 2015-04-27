counter = 0

wrk.method = "POST"

request = function()
  if counter > 9000 then
    counter = 0
  end

  path = "/foo" .. counter .. "/bar" .. counter
  counter = counter + 1
  return wrk.format(nil, path)
end
