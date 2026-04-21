BEGIN;

INSERT INTO accounts (name, type)
VALUES ('Test Account', 'Savings')

-- Example: add expense

-- 1. Insert transaction
INSERT INTO transactions (account_id, amount, description, category, transaction_type)
VALUES (1, -5000, 'Groceries', 7, 2);

-- 2. Update account balance
UPDATE accounts
SET balance = balance + (-5000),
    updated_at = NOW()
WHERE id = 1;

COMMIT;
