const limit = 1_000

function solve(to)
    answer = 0

    for i in 3:to-1
        if i % 3 == 0 || i % 5 == 0
            answer += i
        end
    end

    return answer
end

println(solve(limit))
@time solve(limit)
