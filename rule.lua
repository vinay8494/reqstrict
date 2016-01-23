function filter (request)
  if string.find(request,"google") then
    return false
  else
    return true
  end
end
