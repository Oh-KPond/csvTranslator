require 'csv'

table = CSV.parse(File.read("data/translation_key02192020.csv"), headers: true)

print(table)