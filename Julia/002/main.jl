const limit = 4_000_000

function solve(to)
    answer = 0
    n, m = 1, 1

    while m < to
        m += n
        n = m - n

        if m % 2 == 0
            answer += m
        end
    end

    return answer
end

println(solve(limit))
@time solve(limit)
