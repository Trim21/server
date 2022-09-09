---@type integer
local keyLen = #KEYS

local ttl = ARGV[1]

for i = 1, keyLen do
    local value = redis.call('set', KEYS[i], ARGV[i + 1], 'ex', ttl)
    -- do some logic on the value
end


return keyLen
