readLine = ->
  input_stdin_array[input_currentline++]

main = ->
  a = parseInt(readLine())
  b = parseInt(readLine())
  console.log a+b
  return

process.stdin.resume()
process.stdin.setEncoding 'ascii'
input_stdin = ''
input_stdin_array = ''
input_currentline = 0
process.stdin.on 'data', (data) ->
  input_stdin += data
  return
process.stdin.on 'end', ->
  input_stdin_array = input_stdin.split('\n')
  main()
  return